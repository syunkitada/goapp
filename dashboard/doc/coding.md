# Coding


## Exportの方式
* https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export
* モジュールのexport方法はnamed exportとdefault exportの2種類がある
* 基本的には、default exportを推奨する
* また、import時は、モジュール名をそのまま利用する


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
