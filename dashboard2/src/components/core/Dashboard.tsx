import data from "../../data";
import locationData from "../../data/locationData";

function renderServices(input: any) {
    const { id, idPrefix, serviceName, projectName, onClickService } = input;
    const { ServiceMap, ProjectServiceMap } = data.auth.Authority;

    let tmpServiceMap: any = null;
    let projectText: string;
    if (projectName) {
        projectText = projectName;
        tmpServiceMap = ProjectServiceMap[projectName].ServiceMap;
    } else {
        projectText = "Projects";
        tmpServiceMap = ServiceMap;
    }
    console.log("DEBUG projectName", projectName, tmpServiceMap);

    const tmpProjects = Object.keys(ProjectServiceMap);
    tmpProjects.sort();

    const projectsHtmls = [];
    for (const tmpProject of tmpProjects) {
        projectsHtmls.push(`
        <a class="list-group-item list-group-item-action ${idPrefix}-Project" href="#">${tmpProject}</a>
        `);
    }

    const projectHtml = `
    <li class="list-group-item list-group-item-action sidebar-item">
      <a class="list-group-item-action" data-toggle="collapse" href="#${idPrefix}projects">
        ${projectText}
        <i id="${idPrefix}projectsIcon" class="material-icons">chevron_right</i>
      </a>
      <div class="collapse list-group list-group-flush" id="${idPrefix}projects" style="padding: 5px;">
        <input id="${idPrefix}inputProject" class="form-control form-control-sm">
        ${projectsHtmls.join("")}
      </div>
    </li>
    `;

    const servicesHtmls = [projectHtml];

    const tmpServices = Object.keys(tmpServiceMap);
    tmpServices.sort();

    for (const service of tmpServices) {
        let className = "";
        if (service === serviceName) {
            className = "sidebar-item-active";
        }
        servicesHtmls.push(`
    <li class="list-group-item list-group-item-action sidebar-item">
      <a class="list-group-item-action ${idPrefix}-Service ${className}" href="#">${service}</a>
    </li>
    `);
    }

    $(`#${id}`).html(servicesHtmls.join(""));

    $(`#${idPrefix}projects`)
        .on("show.bs.collapse", function () {
            $(`${idPrefix}projectsIcon`).toggleClass("rotate-90");
        })
        .on("shown.bs.collapse", function () {
            $(`#${idPrefix}inputProject`).focus();
        })
        .on("hide.bs.collapse", function () {
            $(`${idPrefix}projectsIcon`).toggleClass("rotate-90");
        });

    $(`.${idPrefix}-Service`).on("click", function (e) {
        const serviceName = $(this).text();
        onClickService({
            projectName: projectName,
            serviceName: serviceName
        });

        renderServices(Object.assign({}, input, { projectName, serviceName }));
    });

    $(`.${idPrefix}-Project`).on("click", function (e) {
        const projectName = $(this).text();
        const serviceName = "HomeProject";
        onClickService({
            projectName,
            serviceName
        });

        renderServices(Object.assign({}, input, { projectName, serviceName }));
    });
}

function Render(input: any) {
    const { id } = input;
    const { Name } = data.auth.Authority;

    const { serviceName, projectName } = locationData.getServiceParams();

    const idPrefix = `${input.id}-Dashboard-`;

    $(`#${id}`).html(`
    <nav class="navbar navbar-expand-lg navbar-light border-bottom sticky-top bg-white" style="height: 50px; padding: 0px;">
      <a id="menu-toggle" class="border-right" href="#">
        <span class="navbar-toggler-icon"></span>
      </a>

      <a id="navbar-brand" class="navbar-brand border-right mr-auto" href="#">Home</a>

      <div class="dropdown col-auto bg-lignt border-left">
        <button class="btn dropdown-toggle" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          ${Name}
        </button>
        <div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenuButton">
          <a class="dropdown-item" id="DashboardLogout" href="#">Log out</a>
        </div>
      </div>
    </nav>

    <!-- Sidebar -->
    <div class="border-right bg-white" id="sidebar-wrapper">
      <ul id="${idPrefix}-Services" class="list-group list-group-flush" style="width: 100%;">
      </ul>
    </div>
    <!-- /#sidebar-wrapper -->

    <div class="bg-white" id="wrapper">
    <!-- Page Content -->
    <div id="page-content-wrapper">

      <div class="container-fluid">
        <div id="root-content"></div>
      </div>
    </div>
    <!-- /#page-content-wrapper -->

  </div>


  `);

    renderServices(
        Object.assign({}, input, {
            id: `${idPrefix}-Services`,
            idPrefix: idPrefix,
            serviceName,
            projectName
        })
    );

    $("#menu-toggle").on("click", function (e) {
        e.preventDefault();
        $("#sidebar-wrapper").toggleClass("toggled");
    });

    $("#header-menu-toggle").on("click", function (e) {
        e.preventDefault();
        $("#header-menu").toggleClass("toggled");
    });

    $("#DashboardLogout").on("click", function () {
        input.logout();
    });

    return;
}

const index = {
    Render
};
export default index;
