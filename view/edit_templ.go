// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/hossokawa/go-todo-app/internal/models"

func TaskEdit(task *model.Task) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"edit-div\" class=\"h-full w-full pt-64 flex flex-col items-center\"><form class=\"bg-cement m-auto p-10 border-2 border-zinc-700 w-[40vw]\"><h4 class=\"text-white text-3xl text-center mb-8\">Edit task</h4><div class=\"grid grid-cols-2 items-center\"><label class=\"text-white text-2xl mb-3\">Task ID</label><p class=\"text-white text-2xl ml-[-6rem] mb-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(task.Id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/edit.templ`, Line: 10, Col: 62}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><label for=\"name\" class=\"text-white text-2xl mb-3\">Task name</label> <input type=\"text\" name=\"name\" autocomplete=\"off\" class=\"bg-cement w-auto text-white text-xl px-3 py-2 border-2 border-zinc-700 rounded-lg focus:outline-none focus:border-zinc-500 ml-[-6rem] mb-3\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(task.Name))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <label for=\"completed\" class=\"text-white text-2xl mb-3\">Completed</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if task.Completed {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"checkbox\" name=\"completed\" class=\"w-5 h-5 ml-[-6rem] mb-3\" checked>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"checkbox\" name=\"completed\" class=\"w-5 h-5 ml-[-6rem] mb-3\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><span class=\"flex flex-col justify-center items-center\"><div id=\"error-div\"></div><button type=\"submit\" hx-patch=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/tasks/" + task.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#edit-div\" hx-target-error=\"#error-div\" hx-swap=\"outerHTML\" class=\"text-white text-2xl bg-violet-500 w-full py-1 mt-4 rounded-lg hover:bg-violet-700 transition-colors ease-in-out\">Edit</button></span></form><button type=\"button\" hx-get=\"/\" hx-trigger=\"click\" hx-target=\"#edit-div\" hx-swap=\"outerHTML\" class=\"text-white text-2xl bg-[#4B0C55] w-1/4 py-1 mt-12 rounded-lg hover:bg-[#5C1068] transition-colors ease-in-out\">Back to home</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
