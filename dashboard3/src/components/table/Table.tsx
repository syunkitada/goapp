import data from "../../data";

export function Render(input: any) {
    const { id, View } = input;
    const idPrefix = `${id}-Table-`;
    const tableData = data.service.data[View.DataKey];
    console.log("DEBUG Panels.Render", input, View, tableData);

    $(`#${id}`).html(`
    <div class="row table-wrapper">
      <div class="col s3">
        <div class="input-field">
          <input placeholder="Search" type="text">
        </div>
      </div>
      <div class="col s3">
      </div>
      <div class="col s6 pagenation-wrapper">
        <ul class="pagination right">
          <li>
            Rows per page:
            <div class="input-field inline">
             <select class="browser-default">
               <option value="1">100</option>
               <option value="2">200</option>
               <option value="3">300</option>
             </select>
            </div>
            1-64 of 68
          </li>
          <li class="disabled"><a href="#!"><i class="material-icons">first_page</i></a></li>
          <li class="disabled"><a href="#!"><i class="material-icons">chevron_left</i></a></li>
          <li class="active"><a href="#!">1</a></li>
          <li class="waves-effect"><a href="#!">2</a></li>
          <li class="waves-effect"><a href="#!">3</a></li>
          <li class="waves-effect"><a href="#!">4</a></li>
          <li class="waves-effect"><a href="#!">5</a></li>
          <li class="waves-effect"><a href="#!"><i class="material-icons">chevron_right</i></a></li>
          <li class="waves-effect"><a href="#!"><i class="material-icons">last_page</i></a></li>
        </ul>
      </div>

      <div class="col s12" style="overflow: auto;">
        <table class="table">
          <thead><tr id="${idPrefix}thead"></tr></thead>
          <tbody id="${idPrefix}tbody"></tbody>
        </table>
      </div>
    </div>
    `);

    const columns = View.Columns;
    const thHtmls: any = [];
    const isSelectActions = true;
    thHtmls.push(
        `<th class="checkbox-wrapper"><label><input type="checkbox" /><span></span></label></th>`
    );
    for (let i = 0, len = columns.length; i < len; i++) {
        const column = columns[i];
        let alignClass = "right-align";
        if (i === 0) {
            alignClass = "left-align";
        }
        if (column.Align) {
            alignClass = "left-align";
        }
        thHtmls.push(`<th class="${alignClass}">${column.Name}</th>`);
    }
    $(`#${idPrefix}thead`).html(thHtmls.join(""));

    const linkClass = `${idPrefix}tableLink`;
    const bodyTrHtmls: any = [];
    for (let i = 0, len = tableData.length; i < len; i++) {
        const rdata = tableData[i];
        const tdHtmls: any = ["<tr>"];
        if (isSelectActions) {
            tdHtmls.push(
                `<td class="checkbox-wrapper"><label><input type="checkbox" /><span></span></label></td>`
            );
        }

        for (let j = 0, lenj = columns.length; j < lenj; j++) {
            const column = columns[j];
            let alignClass = "right-align";
            if (j === 0) {
                alignClass = "left-align";
            }
            if (column.Align) {
                alignClass = "left-align";
            }

            if (column.LinkPath) {
                tdHtmls.push(
                    `<td class="link ${alignClass} ${linkClass}" id="${idPrefix}${i}-${j}">
                ${rdata[column.Name]}
                </td>`
                );
            } else {
                tdHtmls.push(
                    `<td class="${alignClass}" id="${idPrefix}${i}-${j}">
                ${rdata[column.Name]}
                </td>`
                );
            }
        }

        tdHtmls.push("</tr>");
        $.merge(bodyTrHtmls, tdHtmls);
    }

    console.log("DEBUG", bodyTrHtmls);
    $(`#${idPrefix}tbody`).html(bodyTrHtmls.join(""));

    $(`.${linkClass}`).on("click", function () {
        const id = $(this).attr("id");
        if (id) {
            const splitedId = id.split("-");
            const column = columns[splitedId[splitedId.length - 1]];
            const rdata = tableData[splitedId[splitedId.length - 2]];
            const params: any = {};
            params[column.LinkKey] = rdata[column.Name];
            const location = {
                Path: column.LinkPath,
                Params: params,
                SearchQueries: {}
            };

            console.log("DEBUG click link", column, rdata, location);
        }
    });
}

const index = {
    Render
};
export default index;
