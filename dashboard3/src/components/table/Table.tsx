import data from "../../data";

export function Render(input: any) {
	const { id, View } = input;
	console.log("DEBUG Panels.Render", input, View, data);

	$(`#${id}`).html(`
<form>
  <div class="form-row">
    <div class="col-md-6 mb-3">
      <input type="text" class="form-control">
    </div>
    <div class="col-md-6 mb-3">
      <label class="col-form-label">
        <a><i class="material-icons">first_page</i></a>
        <a><i class="material-icons">chevron_left</i></a>
		<a class="btn">1</a>
        <a><i class="material-icons">chevron_right</i></a>
        <a><i class="material-icons">last_page</i></a>
      </label>
    </div>
  </div>
</form>
    <table class="table table-sm">

  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">

  <div class="form-row">
      <label class="col-form-label">
        <a><i class="material-icons">first_page</i></a>
        <a><i class="material-icons">chevron_left</i></a>
        <a><i class="material-icons">chevron_right</i></a>
        <a><i class="material-icons">last_page</i></a>
      </label>
  </div>

</th>
      <th scope="col">Last</th>
      <th scope="col">Handle</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">1</th>
      <td>Mark</td>
      <td>Otto</td>
      <td>@mdo</td>
    </tr>
    <tr>
      <th scope="row">2</th>
      <td>Jacob</td>
      <td>Thornton</td>
      <td>@fat</td>
    </tr>
    <tr>
      <th scope="row">3</th>
      <td colspan="2">Larry the Bird</td>
      <td>@twitter</td>
    </tr>
  </tbody>

    </div>
    `);
}

const index = {
	Render
};
export default index;
