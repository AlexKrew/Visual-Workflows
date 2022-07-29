class DateTime {
  static getTime(date: Date): string {
    return `${date.getHours()}:${date.getMinutes()}.${date.getSeconds()}`;
  }
}

export default DateTime;
