document.getElementById("add-todo-form").addEventListener("htmx:afterRequest", function(event) {
	if (event.detail.successful) {
		this.reset();
	}
});
