package fcnv

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"firebase.google.com/go/messaging"
)

func TestStruct2FlatMap_1(t *testing.T) {
	type A struct {
		A1 string  `json:"a1"`
		A2 *string `json:"a2"`
	}

	type X struct {
		S    string   `json:"s"`
		SP   *string  `json:"sp"`
		I    int      `json:"i"`
		IP   *int     `json:"ip"`
		I8   int8     `json:"i8"`
		I8P  *int8    `json:"i8p"`
		UI   uint     `json:"ui"`
		UIP  *uint    `json:"uip"`
		UI8  uint8    `json:"ui8"`
		UI8P *uint8   `json:"ui8p"`
		F64  float64  `json:"f64"`
		F64P *float64 `json:"f64p"`
		Bo   bool     `json:"bo"`
		BoP  *bool    `json:"bop"`
		By   []byte   `json:"by"`
		SS   []string
		SPS  []*string
		SA   [2]string
		SPA  [2]*string
		IS   []int
		IPS  []*int
		IA   [2]int
		IPA  [2]*int
		MSS  map[string]string
		MII  map[int]int
		IF   interface{}
		IFS  []interface{}
		ISA  [2]interface{}
		ST   A
		STP  *A
		T    time.Time
		TP   *time.Time
	}

	keyDic := map[string]string{
		"SS":  "ss",
		"SPA": "spa",
		"MSS": "mss",
		"STP": "stp",
		"T":   "time",
	}

	var (
		s   = "test string"
		i   = int(12)
		i8  = int8(23)
		ui  = uint(34)
		ui8 = uint8(45)
		f64 = float64(1.2345)
		bo  = true
		by  = []byte(s)
		now = time.Now().UTC()
	)

	xBefore := X{
		S:    s,
		SP:   &s,
		I:    i,
		IP:   &i,
		I8:   i8,
		I8P:  &i8,
		UI:   ui,
		UIP:  &ui,
		UI8:  ui8,
		UI8P: &ui8,
		F64:  f64,
		F64P: &f64,
		Bo:   bo,
		BoP:  &bo,
		By:   by,
		SS:   []string{s},
		SPS:  []*string{&s},
		SA:   [2]string{s},
		SPA:  [2]*string{&s},
		IS:   []int{i},
		IPS:  []*int{&i},
		IA:   [2]int{i},
		IPA:  [2]*int{&i},
		MSS:  map[string]string{"k1": "v1", "k2": "v2"},
		MII:  map[int]int{1: 2, 2: 3},
		IF:   s,
		IFS:  []interface{}{s, i},
		ISA:  [2]interface{}{s, i},
		ST:   A{A1: s, A2: &s},
		STP:  &A{A1: s, A2: &s},
		T:    now,
		TP:   &now,
	}

	var xAfter X

	if ed, err := Struct2FlatMap(xBefore, keyDic); err != nil {
		t.Error(err)
	} else if err = FlatMap2Struct(ed, &xAfter, keyDic); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(xBefore, xAfter) {
		b1, _ := json.Marshal(xBefore)
		b2, _ := json.Marshal(xAfter)
		if !reflect.DeepEqual(b1, b2) {
			t.Error(string(b1))
			t.Error(string(b2))
		}
	}
}

func TestStruct2FlatMap_2(t *testing.T) {
	now := time.Now().UTC()
	cnt := 3
	an := messaging.AndroidNotification{
		Title:                 "title",
		Body:                  "body",
		Icon:                  "icon/example.png",
		Color:                 "#123456",
		Sound:                 "example/sound",
		Tag:                   "test-tag",
		ClickAction:           "click",
		BodyLocKey:            "body_loc_key",
		BodyLocArgs:           []string{"body_loc_arg1", "body_loc_arg2"},
		TitleLocKey:           "title_loc_key",
		TitleLocArgs:          []string{"title_loc_arg1", "title_loc_arg2"},
		ChannelID:             "default-ch",
		ImageURL:              "image/example.png",
		Ticker:                "ticker",
		Sticky:                true,
		EventTimestamp:        &now,
		LocalOnly:             true,
		Priority:              messaging.PriorityHigh,
		VibrateTimingMillis:   []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		DefaultVibrateTimings: true,
		DefaultSound:          true,
		LightSettings: &messaging.LightSettings{
			Color:                  "#123456",
			LightOnDurationMillis:  1000,
			LightOffDurationMillis: 1000,
		},
		DefaultLightSettings: true,
		Visibility:           messaging.VisibilityPublic,
		NotificationCount:    &cnt,
	}

	keyDic := map[string]string{
		"EventTimestamp":      "event_time",
		"Priority":            "notification_priority",
		"Visibility":          "visibility",
		"VibrateTimings":      "vibrate_timings",
		"VibrateTimingMillis": "vibrate_timings",
	}

	ed, err := Struct2FlatMap(an, keyDic)
	if err != nil {
		t.Error(err)
	}

	an2 := new(messaging.AndroidNotification)
	if err := FlatMap2Struct(ed, an2, keyDic); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(an, *an2) {
		b1, _ := json.Marshal(&an)
		b2, _ := json.Marshal(an2)

		t.Error(string(b1))
		t.Error(string(b2))
	}
}

func TestStruct2FlatMap_3(t *testing.T) {
	type X struct {
		T1 time.Time  `json:"t1"`
		T2 *time.Time `json:"t2"`
	}

	var (
		now = time.Now().UTC()
		x1  = X{T1: now, T2: &now}
		x2  X
	)

	if ed, err := Struct2FlatMap(x1); err != nil {
		t.Error(err)
	} else if err = FlatMap2Struct(ed, &x2); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(x1, x2) {
		b1, _ := json.Marshal(x1)
		b2, _ := json.Marshal(x2)

		t.Error(string(b1))
		t.Error(string(b2))
	}
}
