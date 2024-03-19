// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Login() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"login-div\" class=\"h-screen flex flex-col justify-center items-center\"><form class=\"bg-cement h-1/3 w-1/3 flex flex-col justify-center items-center border border-zinc-700 py-4 group\" novalidate><div for=\"email\" class=\"w-full flex flex-col justify-center items-center\"><label class=\"text-white text-xl\">Email</label> <input type=\"email\" name=\"email\" placeholder=\" \" class=\"peer bg-cement w-3/4 text-white px-3 py-2 border-2 border-zinc-700 rounded-lg focus:outline-none focus:border-zinc-500 invalid:[&amp;:not(:placeholder-shown):not(:focus)]:border-red-500\" required> <span class=\"invisible text-red-500 peer-[&amp;:not(:placeholder-shown):not(:focus):invalid]:visible\">Please enter a valid email address</span></div><div for=\"password\" class=\"w-full flex flex-col justify-center items-center\"><label class=\"text-white text-xl\">Password</label> <input type=\"password\" name=\"password\" autocomplete=\"off\" minlength=\"8\" placeholder=\" \" class=\"peer bg-cement w-3/4 text-white px-3 py-2 border-2 border-zinc-700 rounded-lg focus:outline-none focus:border-zinc-500 invalid:[&amp;:not(:placeholder-shown):not(:focus)]:border-red-500\" required> <span class=\"invisible text-red-500 peer-[&amp;:not(:placeholder-shown):not(:focus):invalid]:visible\">Password must be 8 characters long</span></div><button type=\"submit\" hx-post=\"/login\" hx-trigger=\"click\" hx-target=\"#login-div\" hx-swap=\"outerHTML\" class=\"text-white text-2xl bg-violet-500 w-3/4 py-1 mt-4 rounded-lg hover:bg-violet-700 transition-colors ease-in-out group-invalid:pointer-events-none group-invalid:opacity-30\">Login</button></form><button type=\"button\" hx-get=\"/\" hx-trigger=\"click\" hx-target=\"#login-div\" hx-swap=\"outerHTML\" class=\"text-white text-2xl bg-[#4B0C55] w-auto px-4 py-1 mt-12 rounded-lg hover:bg-[#5C1068] transition-colors ease-in-out \">Back to home</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
