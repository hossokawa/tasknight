// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Input() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"name\" placeholder=\"e.g. code some Go\" autocomplete=\"off\" id=\"todo-input\" class=\"w-full bg-cement text-white text-xl px-5 py-2 border-2 border-zinc-700 rounded-l-lg placeholder:text-zinc-700 placeholder:italic focus:outline-none focus:border-zinc-500 p-2 peer\"> <button type=\"button\" hx-post=\"/\" hx-trigger=\"click\" hx-swap=\"multi:#task-list:outerHTML,#todo-input:outerHTML\" hx-ext=\"multi-swap\" class=\"w-auto text-white text-2xl bg-violet-500 px-8 rounded-r-lg hover:bg-violet-700 transition-colors ease-in-out peer-placeholder-shown:pointer-events-none peer-placeholder-shown:opacity-50\">Submit</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
