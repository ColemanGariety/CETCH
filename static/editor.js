hljs.configure({
  languaes: ["go"]
});

var src = document.querySelector("pre code.editor")
var res = document.querySelector("pre code#result")
hljs.highlightBlock(src)
hljs.highlightBlock(res)

src.addEventListener("keydown", function (event) {
  if (event.keyCode == 9) {
    event.preventDefault();
    document.execCommand('insertHTML', false, '&nbsp;&nbsp;')
  }
})

src.addEventListener("input", function (event) {
  res.innerHTML = src.innerHTML
  hljs.highlightBlock(res)
})
