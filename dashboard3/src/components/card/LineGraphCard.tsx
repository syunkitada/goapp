import Chart from "chart.js";

const colors = [
    ["rgba(255, 99, 132, 0.2)", "rgba(255, 99, 132, 1)"],
    ["rgba(54, 162, 235, 0.2)", "rgba(54, 162, 235, 1)"],
    ["rgba(255, 206, 86, 0.2)", "rgba(255, 206, 86, 1)"],
    ["rgba(75, 192, 192, 0.2)", "rgba(75, 192, 192, 1)"],
    ["rgba(153, 102, 255, 0.2)", "rgba(153, 102, 255, 1)"],
    ["rgba(255, 159, 64, 0.2)", "rgba(255, 159, 64, 1)"]
];

function formatLabel(label: any) {
    if (label < 1000) {
        return label;
    }
    if (label < 1000000) {
        return label / 1000 + "K";
    }
    if (label < 1000000000) {
        return label / 1000000 + "M";
    }
    return label / 1000000000 + "G";
}

export function Render(input: any) {
    const { id, View, metric } = input;

    const keyPrefix = `${id}-LineGraphCard-`;
    const canvasId = `${keyPrefix}canvas`;

    $(`#${id}`).html(`<canvas id="${canvasId}"></canvas>`);

    const elem: any = document.getElementById(canvasId);
    if (!elem) {
        return;
    }
    console.log("DEBUG metric", metric);

    var datasets: any = [];
    for (let i = 0, len = metric.Keys.length; i < len; i++) {
        const data: any = [];
        datasets.push({
            label: metric.Keys[i],
            data,
            pointRadius: 0,
            backgroundColor: colors[i][0],
            borderColor: colors[i][1],
            borderWidth: 1
        });
    }
    for (let i = 0, len = metric.Values.length; i < len; i++) {
        const value = metric.Values[i];
        for (let j = 0, lenj = metric.Keys.length; j < lenj; j++) {
            datasets[j].data.push({
                t: new Date(value["time"]),
                y: value[metric.Keys[j]]
            });
        }
    }
    console.log("DEBUG datasets", datasets);

    var ctx = elem.getContext("2d");
    var myChart = new Chart(ctx, {
        type: "line",
        data: {
            labels: ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
            datasets: datasets
        },
        options: {
            tooltips: {
                mode: "index",
                intersect: false,
                callbacks: {
                    label: function (tooltipItem: any, data: any) {
                        var label =
                            data.datasets[tooltipItem.datasetIndex].label || "";

                        if (label) {
                            label += ": ";
                        }
                        label += formatLabel(tooltipItem.yLabel);
                        return label;
                    }
                }
            },
            hover: {
                mode: "nearest",
                intersect: true
            },
            scales: {
                yAxes: [
                    {
                        ticks: {
                            beginAtZero: true,
                            callback: function (
                                label: any,
                                index: any,
                                labels: any
                            ) {
                                return formatLabel(label);
                            }
                        },
                        scaleLabel: {
                            display: true
                        }
                    }
                ],
                xAxes: [
                    {
                        type: "time",
                        distribution: "series",
                        time: {
                            unit: "minute",
                            tooltipFormat: "MM-DD HH:mm:SS ZZ",
                            parser: "MM/DD",
                            displayFormats: {
                                minute: "MM-DD HH:mm"
                            }
                        }
                    }
                ]
            }
        }
    });
}

const index = {
    Render
};
export default index;
