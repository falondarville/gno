// XXX TEST GITHUB ACTION
// This file is autogenerated from the genstd tool (@/misc/stdgen); do not edit.
// To regenerate it, run `go generate` from this directory.

package stdlibs

import (
	"reflect"

	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	lib0 "github.com/gnolang/gno/gnovm/stdlibs/crypto/sha256"
	lib1 "github.com/gnolang/gno/gnovm/stdlibs/math"
	lib2 "github.com/gnolang/gno/gnovm/stdlibs/std"
	lib3 "github.com/gnolang/gno/gnovm/stdlibs/strconv"
	lib4 "github.com/gnolang/gno/gnovm/stdlibs/time"
)

type nativeFunc struct {
	gnoPkg  string
	gnoFunc gno.Name
	params  []gno.FieldTypeExpr
	results []gno.FieldTypeExpr
	f       func(m *gno.Machine)
}

var nativeFuncs = [...]nativeFunc{
	{
		"crypto/sha256",
		"Sum256",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]byte")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib0.Sum256(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32bits",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint32")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib1.Float32bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32frombits",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float32")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib1.Float32frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64bits",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib1.Float64bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64frombits",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib1.Float64frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"AssertOriginCall",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			lib2.AssertOriginCall(
				m,
			)
		},
	},
	{
		"std",
		"IsOriginCall",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			r0 := lib2.IsOriginCall(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"CurrentRealmPath",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := lib2.CurrentRealmPath(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"GetChainID",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := lib2.GetChainID(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"GetHeight",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0 := lib2.GetHeight(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"Itoa",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib3.Itoa(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"AppendUint",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
			{Name: gno.N("p1"), Type: gno.X("uint64")},
			{Name: gno.N("p2"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]byte")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  uint64
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  int
				rp2 = reflect.ValueOf(&p2).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)

			r0 := lib3.AppendUint(p0, p1, p2)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"Atoi",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int")},
			{Name: gno.N("r1"), Type: gno.X("error")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1 := lib3.Atoi(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"strconv",
		"CanBackquote",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib3.CanBackquote(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"FormatInt",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int64")},
			{Name: gno.N("p1"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int64
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  int
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := lib3.FormatInt(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"FormatUint",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint64")},
			{Name: gno.N("p1"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint64
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  int
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := lib3.FormatUint(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"Quote",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib3.Quote(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"QuoteToASCII",

		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := lib3.QuoteToASCII(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"time",
		"now",

		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
			{Name: gno.N("r1"), Type: gno.X("int32")},
			{Name: gno.N("r2"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0, r1, r2 := lib4.X_now(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r2).Elem(),
			))
		},
	},
}
