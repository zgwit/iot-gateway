import { NzTableQueryParams } from "ng-zorro-antd/table";

export function ParseTableQuery(query: NzTableQueryParams, body: any): void {
  // const body: any = {
  //   filter: {},
  //   //sort: {},
  // }
  if (typeof body.filter === 'undefined') {
    body.filter = {};
  }
  //过滤器
  query.filter.forEach(f => {
    if (f.value.length > 1)
      body.filter[f.key] = f.value;
    else if (f.value.length === 1)
      body.filter[f.key] = f.value[0];
  })

  //分布
  body.skip = (query.pageIndex - 1) * query.pageSize;
  body.limit = query.pageSize;

  //排序
  const sorts = query.sort.filter(s => s.value);
  if (sorts.length) {
    body.sort = {};
    sorts.forEach(s => {
      body.sort[s.key] = s.value === 'ascend' ? 1 : -1;
    });
  } else {
    delete body.sort;
  }
}
