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

			@media (min-width: 768px) {
				.bd-placeholder-img-lg {
				font-size: 3.5rem;
				}
			}
			</style>
			<!-- Custom styles for this template -->
		<link href="sticky-footer.css" rel="stylesheet">
	</head>

	<body class="d-flex flex-column h-100">
		<main>
			<div class="container py-4">
				<header class="pb-3 mb-4 border-bottom">
					<a href="https://github.com/WebAssembly" class="d-flex align-items-center text-dark text-decoration-none">
                    <img src="https://www.vectorlogo.zone/logos/webassembly/webassembly-icon.svg" width="36" height="32" alt="WASM">
					<span class="fs-4">&nbsp;WASM&nbsp;/&nbsp;WASI</span>
					</a>
				</header>

				<div class="p-5 mb-4 bg-light rounded-3">
					<div class="container-fluid py-5">
						<h1 class="display-5 fw-bold">What is WebAssembly?</h1>
						<p class="col-md-8 fs-4">WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications.</p>
						<a href="https://webassembly.org/">
						<button class="btn btn-primary btn-lg" type="button">Official Site</button>
						</a>
					</div>
				</div>

				<div class="row align-items-md-stretch">
					<div class="col-md-6">
						<div class="h-100 p-5 text-bg-dark rounded-3">
							<h2>WebAssembly System Interface</h2>
							<p>The WebAssembly System Interface is not a monolithic standard system interface, but is instead a modular collection of standardized APIs. None of the APIs are required to be implemented to have a compliant runtime. Instead, host environments can choose which APIs make sense for their use cases.</p>
							<a href="https://github.com/WebAssembly/WASI">
							<button class="btn btn-outline-light" type="button">Go Github</button>
							</a>
						</div>
					</div>

					<div class="col-md-6">
						<div class="h-100 p-5 bg-light border rounded-3">
							<h2>Practice</h2>
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
