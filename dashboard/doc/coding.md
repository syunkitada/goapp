# Coding

## Export の方式

- https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export
- モジュールの export 方法は named export と default export の 2 種類がある
- 基本的には、default export を推奨する
- また、import 時は、モジュール名をそのまま利用する

#### Export Example

lib/sort_utils/index.tsx

```
import sortNums from './sortNums';
import sortStrs from './sortStrs';

export default {
  sortNums,
  sortStrs,
}
```

lib/sort_utils/sortNums.tsx

```
export default function sortNums(...) {
  ...
}
```

lib/sort_utils/sortNums.test.tsx

```
it('sortNums', () => {
  ...
}
```

lib/sort_utils/sortStrs.tsx

```
export default function sortStrs(...) {
  ...
}
```

lib/sort_utils/sortStrs.test.tsx

```
it('sortStr', () => {
  ...
}
```

#### Import Example

```
import sort_utils from "../../lib/sort_utils";

sort_utils.sortNums(...)
```
