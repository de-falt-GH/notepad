const relativeTimeCutoffsMs = [
	60,
	3600,
	86400,
	86400 * 7,
	86400 * 30,
	86400 * 365,
	Infinity,
]

const timeUnits: Intl.RelativeTimeFormatUnit[] = [
	'second',
	'minute',
	'hour',
	'day',
	'week',
	'month',
	'year',
]

type GetRelativeTimeOptions = Readonly<{
	originDate?: Date
	locales?: Intl.LocalesArgument
}>

function dateTimeMs(date: Date | number) {
	return typeof date == 'number' ? date : date.getTime()
}

export function getRelativeTimeString(
	date: Date | number,
	options?: GetRelativeTimeOptions,
): string {
	const nowTimeMs = dateTimeMs(date)
	const wasTimeMs = dateTimeMs(options?.originDate ?? Date.now())

	const dtSec = Math.round((nowTimeMs - wasTimeMs) / 1000)

	const unitIndex = relativeTimeCutoffsMs.findIndex(
		cutoff => cutoff > Math.abs(dtSec),
	)

	const divisor = unitIndex ? relativeTimeCutoffsMs[unitIndex - 1] : 1

	const rtf = new Intl.RelativeTimeFormat(options?.locales, {
		numeric: 'auto',
	})
	return rtf.format(Math.floor(dtSec / divisor), timeUnits[unitIndex])
}
