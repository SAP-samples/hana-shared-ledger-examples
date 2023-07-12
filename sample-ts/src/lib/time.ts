import moment from "moment"

const timeLayout = "YYYY-MM-DD HH:mm:ss.SSSSSSS"

export function currentTime(): string {
    return moment.utc().format(timeLayout)
}

