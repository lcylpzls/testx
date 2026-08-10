package core

import (
	"reflect"
	"strings"
)

// Contains 断言容器包含元素：
// 字符串检查子串，切片/数组检查元素，map 检查键。
func Contains(t TB, container, elem any) {
	t.Helper()
	if s, ok := container.(string); ok {
		sub, ok := elem.(string)
		if !ok {
			t.Fatalf("Contains 的字符串容器要求元素为 string，得到 %T", elem)
			return
		}
		if !strings.Contains(s, sub) {
			t.Errorf("期望 %q 包含 %q", s, sub)
		}
		return
	}
	rv := reflect.ValueOf(container)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if reflect.DeepEqual(rv.Index(i).Interface(), elem) {
				return
			}
		}
		t.Errorf("期望容器 %s 包含 %s", formatValue(container), formatValue(elem))
	case reflect.Map:
		ev := reflect.ValueOf(elem)
		if ev.IsValid() && ev.Type().AssignableTo(rv.Type().Key()) {
			if v := rv.MapIndex(ev); v.IsValid() {
				return
			}
		}
		t.Errorf("期望容器 %s 包含键 %s", formatValue(container), formatValue(elem))
	default:
		t.Fatalf("Contains 不支持容器类型 %T", container)
	}
}

// NotContains 断言容器不包含元素；语义与 Contains 相同。
func NotContains(t TB, container, elem any) {
	t.Helper()
	if s, ok := container.(string); ok {
		sub, ok := elem.(string)
		if !ok {
			t.Fatalf("NotContains 的字符串容器要求元素为 string，得到 %T", elem)
			return
		}
		if strings.Contains(s, sub) {
			t.Errorf("期望 %q 不包含 %q", s, sub)
		}
		return
	}
	rv := reflect.ValueOf(container)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if reflect.DeepEqual(rv.Index(i).Interface(), elem) {
				t.Errorf("期望容器 %s 不包含 %s", formatValue(container), formatValue(elem))
				return
			}
		}
	case reflect.Map:
		ev := reflect.ValueOf(elem)
		if ev.IsValid() && ev.Type().AssignableTo(rv.Type().Key()) {
			if v := rv.MapIndex(ev); v.IsValid() {
				t.Errorf("期望容器 %s 不包含键 %s", formatValue(container), formatValue(elem))
				return
			}
		}
	default:
		t.Fatalf("NotContains 不支持容器类型 %T", container)
	}
}

// Subset 断言 list 包含 sublist 的全部元素（考虑多重性，顺序无关）。
func Subset(t TB, list, sublist any) {
	t.Helper()
	lv := reflect.ValueOf(list)
	sv := reflect.ValueOf(sublist)
	if !isListKind(lv) || !isListKind(sv) {
		t.Fatalf("Subset 仅支持切片/数组，得到 %T 与 %T", list, sublist)
		return
	}
	if !subElementsMatch(lv, sv) {
		t.Errorf("期望列表 %s 包含子集 %s", formatValue(list), formatValue(sublist))
	}
}

// ElementsMatch 断言两个集合元素一致（顺序无关、多重性一致）。
func ElementsMatch(t TB, listA, listB any) {
	t.Helper()
	av := reflect.ValueOf(listA)
	bv := reflect.ValueOf(listB)
	if !isListKind(av) || !isListKind(bv) {
		t.Fatalf("ElementsMatch 仅支持切片/数组，得到 %T 与 %T", listA, listB)
		return
	}
	if av.Len() != bv.Len() || !subElementsMatch(bv, av) {
		t.Errorf("期望元素一致（顺序无关）：\n  A: %s\n  B: %s",
			formatValue(listA), formatValue(listB))
	}
}

// isListKind 判断反射值是否为切片/数组。
func isListKind(rv reflect.Value) bool {
	if !rv.IsValid() {
		return false
	}
	k := rv.Kind()
	return k == reflect.Slice || k == reflect.Array
}

// subElementsMatch 判断 sub 的每个元素都能在 list 中找到未使用的相等元素。
func subElementsMatch(list, sub reflect.Value) bool {
	used := make([]bool, list.Len())
	for i := 0; i < sub.Len(); i++ {
		found := false
		for j := 0; j < list.Len(); j++ {
			if used[j] {
				continue
			}
			if reflect.DeepEqual(list.Index(j).Interface(), sub.Index(i).Interface()) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
