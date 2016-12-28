var source;

document.querySelectorAll("[draggable='true']").forEach(function (elem) {
  elem.addEventListener('dragstart', function (event) {
    source = elem;
    event.dataTransfer.setData("text/plain", this.getAttribute("value"))
  })
})

document.querySelectorAll(".reorderable-target").forEach(function (elem) {
  elem.addEventListener('dragover', function (event) {
    event.preventDefault()
    event.dataTransfer.dropEffect = "copy"
  })
  elem.addEventListener('drop', function (event) {
    event.preventDefault()
    console.log(event)
    var temp = source.value
    source.value = event.target.value
    event.target.value = temp
  })
})
