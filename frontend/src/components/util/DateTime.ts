class DateTime {
  static getTime(date: Date): string {
    const hours = DateTime.getLeadingZero(date.getHours());
    const minutes = DateTime.getLeadingZero(date.getMinutes());
    const seconds = DateTime.getLeadingZero(date.getSeconds());
    return `${hours}:${minutes}.${seconds}`;
  }

  static getLeadingZero(time: number): string {
    return ("0" + time.toString()).slice(-2);
  }
}

export default DateTime;
