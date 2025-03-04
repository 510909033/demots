package usecase

import (
	"testing"
	"time"
)

func TestGetTodayStartTs(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	result := GetTodayStartTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestGetTodayEndTs(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
	result := GetTodayEndTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestGetWeekStartTs(t *testing.T) {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	expected := time.Date(now.Year(), now.Month(), now.Day()-weekday+1, 0, 0, 0, 0, now.Location()).Unix()
	result := GetWeekStartTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestGetWeekEndTs(t *testing.T) {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	expected := time.Date(now.Year(), now.Month(), now.Day()-weekday+7, 23, 59, 59, 0, now.Location()).Unix()
	result := GetWeekEndTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestGetMonthStartTs(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
	result := GetMonthStartTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestGetMonthEndTs(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location()).Unix()
	result := GetMonthEndTs()
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
