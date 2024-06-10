export const debounce = <Args extends unknown[]>(
	func: (...args: Args) => void,
	wait = 500,
) => {
	let timeout: ReturnType<typeof setTimeout>

	return (...args: Args) => {
		clearTimeout(timeout)
		timeout = setTimeout(() => func(...args), wait)
	}
}
