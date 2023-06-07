package index

var (
	index string = `
	<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<meta name="description" content="">
		<meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
		<meta name="generator" content="Jekyll v3.8.5">

		<title>Example</title>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-rbsA2VBKQhggwzxH7pPCaAqO46MgnOM80zW1RWuH61DGLwZJEdK2Kadq2F9CUG65" crossorigin="anonymous">

			<style>
				.bd-placeholder-img {
					font-size: 1.125rem;
					text-anchor: middle;
					-webkit-user-select: none;
					-moz-user-select: none;
					-ms-user-select: none;
					user-select: none;
				}

				.carousel-control-next, .carousel-control-prev {
					width: 5%;
				}

				.carousel-indicators {
					margin-bottom: 2rem;
				}

				.container-fluid, .py-5 {
					padding-left: calc(var(--bs-gutter-x) * .99)
				}

				@media (min-width: 768px) {
					.bd-placeholder-img-lg {
					font-size: 3.5rem;
					}
				}
			</style>
	</head>

	<body class="d-flex flex-column h-100">
		<main>
			<div class="container py-4">
				<header class="pb-3 mb-4 border-bottom">
					<a href="https://github.com/WebAssembly" class="d-flex align-items-center text-dark text-decoration-none">
                    <img src="https://www.vectorlogo.zone/logos/webassembly/webassembly-icon.svg" width="36" height="32" alt="WASM">
					<span class="fs-4">&nbsp;WebAssembly</span>
					</a>
				</header>

				<div id="carouselExampleIndicators" class="carousel carousel-dark slide" data-bs-ride="false" data-bs-interval="false">

					<div class="carousel-indicators">
						<button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
						<button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1" aria-label="Slide 2"></button>
						<button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2" aria-label="Slide 3"></button>
						<button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="3" aria-label="Slide 4"></button>
						<button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="4" aria-label="Slide 5"></button>
					</div>

					<div class="carousel-inner">
						<div class="carousel-item active">
							<div class="p-5 mb-4 bg-light rounded-3">
								<div class="container-fluid py-5">
									<h1 class="display-5 fw-bold">What is WebAssembly?</h1>
									<p class="col-md-9 fs-4">WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications.</p>
									<a href="https://webassembly.org/">
										<button class="btn btn-primary btn-lg" type="button">Official Site</button>
									</a>
								</div>
							</div>
						</div>

						<div class="carousel-item">
							<div class="p-5 mb-4 bg-light rounded-3">
								<div class="container-fluid py-5">
									<h1 class="display-20 fw-bold">WebAssembly And JavaScript</h1>
									<br/>
									<p class="col-md-10 fs-6"> 
										* WebAssembly(WASM) 실제로 네이티브 코드를 웹에서 실행하게 해주는 도구 <br/><br/>
										* 보통 브라우저 Runtime의 경우 (JavaScript와 같이) JavaScript 엔진 내에서 나란히 실행 <br/><br/>
										* WASM은 JavaScript에 비해 빠르게 실행되는 구조로, 일반적으로 더 '빠르게' 실행 <br/><br/>
										&nbsp;&nbsp;&nbsp;- 참고 : <a href="https://hacks.mozilla.org/2017/02/what-makes-webassembly-fast/">What makes WebAssembly fast?</a>
									</p>
								</div>
							</div>
						</div>

						<div class="carousel-item">
							<div class="p-5 mb-4 bg-light rounded-3">
								<div class="container-fluid py-5">
									<h1 class="display-20 fw-bold">A Real-World WebAssembly Benchmark MacOS</h1>
									<canvas class="my-4 w-100 chartjs-render-monitor" id="myChart" width="3000" height="489" style="height: 500px; width: 1552px;"></canvas>
									<script src="https://cdn.jsdelivr.net/npm/feather-icons@4.28.0/dist/feather.min.js" integrity="sha384-uO3SXW5IuS1ZpFPKugNNWqTZRRglnUJK6UAZ/gxOX80nxEkN9NcGZTftn6RzhGWE" crossorigin="anonymous"></script>
									<script src="https://cdn.jsdelivr.net/npm/chart.js@2.9.4/dist/Chart.min.js" integrity="sha384-zNy6FEbO50N+Cg5wap8IKA4M/ZnLJgzc6w2NqACZaK0u0FXfOWRRJOnQtpZun8ha" crossorigin="anonymous"></script>
									<script src="http://localhost:3000/dashboard.js"></script>
									문서를 PDF로 변환하는 도구 - PSPDFKit의 벤치마킹 자료 <br/>
									<a href="https://pspdfkit.com/blog/2018/a-real-world-webassembly-benchmark">https://pspdfkit.com/blog/2018/a-real-world-webassembly-benchmark</a>
								</div>
							</div>
						</div>

						<div class="carousel-item">
							<div class="p-5 mb-4 bg-light rounded-3">
								<div class="container-fluid py-5">
									<h1 class="display-8 fw-bold">Movement for the New WebAssembly Standard, WASI</h1>
									<br/>
									Mozilla - 2019 / 3 / 27 발표자료 <br/>
									<a href="https://hacks.mozilla.org/2019/03/standardizing-wasi-a-webassembly-system-interface/">Standardizing WASI: A system interface to run WebAssembly outside the web</a>	
									<br/>
								</div>
							</div>
						</div>

						<div class="carousel-item">
							<div class="p-5 mb-4 bg-light rounded-3">
							<h1 class="display-8 fw-bold">WebAssembly Pipeline</h1>
								<div class="container-fluid py-5 d-flex justify-content-center">
									<div>
										<img src="https://cdn.jsdelivr.net/gh/b0xt/sobyte-images1/2023/03/28/2130b6d3eed84a34b97c6e43d63c2f5e.png" width="750" height="320" alt="Support">
									</div>
								</div>
							</div>
						</div>

						<button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="prev">
							<span class="carousel-control-prev-icon" aria-hidden="false"></span>
							<span class="visually-hidden">Previous</span>
						</button>
						<button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="next">
							<span class="carousel-control-next-icon" aria-hidden="false"></span>
							<span class="visually-hidden">Next</span>
						</button>

					</div>
				</div>

				<div class="row align-items-md-stretch">
					<div class="col-md-6">
						<div class="h-100 p-5 text-bg-dark rounded-3">
							<img src="https://wasi.dev/polyfill/WASI-small.png" width="70" height="50" alt="WASI">
							<h2>WebAssembly System Interface</h2>
							<p>The WebAssembly System Interface is not a monolithic standard system interface, but is instead a modular collection of standardized APIs. None of the APIs are required to be implemented to have a compliant runtime. Instead, host environments can choose which APIs make sense for their use cases.</p>
							<a href="https://github.com/WebAssembly/WASI">
							<button class="btn btn-outline-light" type="button">Go Github</button>
							</a>
						</div>
					</div>

					<div class="col-md-6">
						<div class="h-100 p-5 bg-light border rounded-3">
							<h2>Demo</h2>
							<br/>
							<p>Reference
							<br/>
								&nbsp;"2020년과 이후 JavaScript의 동향"
								<br/>
								- WebAssembly Naver D2 Blog <a href="https://d2.naver.com/helloworld/8257914">https://d2.naver.com/helloworld/8257914</a>
							</p>
							<br/>
							<p>Let's start with a demo of WASM/WASI!</p>
							<a href="http://localhost:3000/poc">
								<button class="btn btn-outline-secondary" type="button">Go!</button>
							</a>
						</div>
					</div>
				</div>

			</div>
		</main>
	</body>
	</html>
	`
)

func GetIndex() string {
	return index
}
