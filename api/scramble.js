function scrambler(string) {
	let characters = string.split("")
	for (var i = 0; i < characters.length; i++) {
		var span = $('<span />').addClass('hidden').text(characters[i]);
		$('body').append(span)
	}
}