const sse = new EventSource('http://localhost:3000/notify')

sse.addEventListener('error', (event) => {
	if (event.readyState === EventSource.CLOSED) {
		console.log('Connection was closed')
		return
	}

	console.log('Error occurred', event)
})

sse.addEventListener('open', (event) => {
	console.log('Connection was opened')
})

sse.addEventListener('new_message', (event) => {
	console.log('New message', event)
})

sse.addEventListener('delete_message', (event) => {
	console.log('Delete message', event)
})
