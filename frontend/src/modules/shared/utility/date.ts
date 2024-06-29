import { format as tFormat, 
    tzDate as ttzDate,
    addDay as tAddDay,
} from "@formkit/tempo"

export const getDate = (date: Date, format='YYYY-MM-DD', timeZone = 'America/New_York') => {
    return tFormat(tzDate(date, timeZone),format);
}

export const addDay = (date: Date, daysQty: number): Date => {
    return tAddDay(date, daysQty);
}

export const tzDate = (date: Date, timeZone = 'America/New_York'): Date => {
    return ttzDate(date, timeZone);
}

export const getLastMonday = (date: Date, timeZone = 'America/New_York'): Date => {
    const currentDate = ttzDate(date, timeZone);
    // getting last monday from current date
    return new Date(currentDate.setDate(currentDate.getDate() - currentDate.getDay() + (currentDate.getDay() === 0 ? -6 : 1)));
}