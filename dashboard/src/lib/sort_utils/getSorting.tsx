function desc(a: any[], b: any[], orderBy: number) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

export default function getSorting(order: string, orderBy: number) {
  return order === 'desc'
    ? (a: any[], b: any[]) => desc(a, b, orderBy)
    : (a: any[], b: any[]) => -desc(a, b, orderBy);
}
