package index

var (
	dashboard string = `
			/* globals Chart:false, feather:false */

			(() => {
			"use strict";

			feather.replace({ "aria-hidden": "true" });

			const opt = {
				events: false,
				tooltips: {
				enabled: false,
				},
				hover: {
				animationDuration: 0,
				},
				animation: {
				duration: 1,
				onComplete: function () {
					var chartInstance = this.chart,
					ctx = chartInstance.ctx;
					ctx.font = Chart.helpers.fontString(
					Chart.defaults.global.defaultFontSize,
					Chart.defaults.global.defaultFontStyle,
					Chart.defaults.global.defaultFontFamily
					);
					ctx.textAlign = "center";
					ctx.textBaseline = "bottom";

					this.data.datasets.forEach(function (dataset, i) {
					var meta = chartInstance.controller.getDatasetMeta(i);
					meta.data.forEach(function (bar, index) {
						var data = dataset.data[index];
						ctx.fillText(data, bar._model.x, bar._model.y - 5);
					});
					});
				},
				},
			};

			// Graphs
			const ctx = document.getElementById("myChart");
			// eslint-disable-next-line no-unused-vars
			const myChart = new Chart(ctx, {
				type: "bar",
				data: {
				labels: [
					"Chrome 67.0.3396 64-bit",
					"Chrome 69.0.3481 64-bit new Baseline Compiler",
					"Firefox 61.0 64-bit",
					"Safari 11.1.1 13605.2.8",
				],
				datasets: [
					{
					label: "WebAssembly",
					data: [5408, 4325, 1902, 8382],
					lineTension: 0,
					backgroundColor: "rgba(255, 99, 132, 0.2)",
					borderColor: "rgba(255, 99, 132, 1)",
					borderWidth: 1,
					},
					{
					label: "JavaScript fallback",
					data: [4784, 4797, 6855, 5802],
					lineTension: 0,
					backgroundColor: "rgba(54, 162, 235, 0.2)",
					borderColor: "rgba(54, 162, 235, 1)",
					borderWidth: 1,
					},
				],
				},
				options: opt,
			});
			})();
		`
)

func GetDashboard() string {
	return dashboard
}
