package strings

import (
    "reflect"
    "testing"
)

func TestStringCalculator_Add(t *testing.T) {
    type args struct {
        numbers string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "empty string returns 0",
            args: args{numbers: ""},
            want: 0,
        },
        {
            name: "a single number returns the value",
            args: args{numbers: "321"},
            want: 321,
        },
        {
            name: "two comma delimited numbers return sum",
            args: args{numbers: "11,7"},
            want: 18,
        },
        {
            name: "two newline delimited numbers return sum",
            args: args{numbers: "10\n2"},
            want: 12,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sc := &StringCalculator{}
            if got := sc.Add(tt.args.numbers); got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestStringCalculator_splitTokens(t *testing.T) {
    type args struct {
        text string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name: "two tokens separated by coma should be split",
            args: args{text: "a,b"},
            want: []string{"a", "b"},
        },
        {
            name: "two tokens separated by space should be split",
            args: args{text: "a\nb"},
            want: []string{"a", "b"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sc := &StringCalculator{}
            if got := sc.splitTokens(tt.args.text); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("splitTokens() = %v, want %v", got, tt.want)
            }
        })
    }
}