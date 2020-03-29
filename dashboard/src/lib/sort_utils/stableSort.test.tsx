import stableSort from "./stableSort";
import getSorting from "./getSorting";


it('getSorting', () => {
    const items = [[1, 5], [3, 4], [2, 3]];
    var sortedItems: any[][];
    var expectedSortedItems: any[][];

    sortedItems = stableSort(items, getSorting("", 0))
    expectedSortedItems = [[1, 5], [2, 3], [3, 4]];
    expect(sortedItems).toEqual(expectedSortedItems)

    sortedItems = stableSort(items, getSorting("desc", 0))
    expectedSortedItems = [[3, 4], [2, 3], [1, 5]];
    expect(sortedItems).toEqual(expectedSortedItems)

    sortedItems = stableSort(items, getSorting("", 1))
    expectedSortedItems = [[2, 3], [3, 4], [1, 5]];
    expect(sortedItems).toEqual(expectedSortedItems)

    sortedItems = stableSort(items, getSorting("desc", 1))
    expectedSortedItems = [[1, 5], [3, 4], [2, 3]];
    expect(sortedItems).toEqual(expectedSortedItems)
});
