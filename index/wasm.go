package index

var (
	wasm string = `
	<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<title>Example</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-rbsA2VBKQhggwzxH7pPCaAqO46MgnOM80zW1RWuH61DGLwZJEdK2Kadq2F9CUG65" crossorigin="anonymous">

			<style>
				.bd-placeholder-img {
				font-size: 1.125rem;
				text-anchor: middle;
				-webkit-user-select: none;
				-moz-user-select: none;
				user-select: none;
				}
		
				@media (min-width: 768px) {
				.bd-placeholder-img-lg {
					font-size: 3.5rem;
				}
				}
		
				.b-example-divider {
				height: 3rem;
				background-color: rgba(0, 0, 0, .1);
				border: solid rgba(0, 0, 0, .15);
				border-width: 1px 0;
				box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
				}
		
				.b-example-vr {
				flex-shrink: 0;
				width: 1.5rem;
				height: 100vh;
				}
		
				.bi {
				vertical-align: -.125em;
				fill: currentColor;
				}
		
				.nav-scroller {
				position: relative;
				z-index: 2;
				height: 2.75rem;
				overflow-y: hidden;
				}
		
				.nav-scroller .nav {
				display: flex;
				flex-wrap: nowrap;
				padding-bottom: 1rem;
				margin-top: -1px;
				overflow-x: auto;
				text-align: center;
				white-space: nowrap;
				-webkit-overflow-scrolling: touch;
				}
			</style>
			<svg xmlns="http://www.w3.org/2000/svg" style="display: none;">
				<symbol id="tools" viewBox="0 0 16 16">
				<path d="M1 0L0 1l2.2 3.081a1 1 0 0 0 .815.419h.07a1 1 0 0 1 .708.293l2.675 2.675-2.617 2.654A3.003 3.003 0 0 0 0 13a3 3 0 1 0 5.878-.851l2.654-2.617.968.968-.305.914a1 1 0 0 0 .242 1.023l3.356 3.356a1 1 0 0 0 1.414 0l1.586-1.586a1 1 0 0 0 0-1.414l-3.356-3.356a1 1 0 0 0-1.023-.242L10.5 9.5l-.96-.96 2.68-2.643A3.005 3.005 0 0 0 16 3c0-.269-.035-.53-.102-.777l-2.14 2.141L12 4l-.364-1.757L13.777.102a3 3 0 0 0-3.675 3.68L7.462 6.46 4.793 3.793a1 1 0 0 1-.293-.707v-.071a1 1 0 0 0-.419-.814L1 0zm9.646 10.646a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708zM3 11l.471.242.529.026.287.445.445.287.026.529L5 13l-.242.471-.026.529-.445.287-.287.445-.529.026L3 15l-.471-.242L2 14.732l-.287-.445L1.268 14l-.026-.529L1 13l.242-.471.026-.529.445-.287.287-.445.529-.026L3 11z"/>
				</symbol>
				<symbol id="cpu-fill" viewBox="0 0 16 16">
					<path d="M6.5 6a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-3z"/>
					<path d="M5.5.5a.5.5 0 0 0-1 0V2A2.5 2.5 0 0 0 2 4.5H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2A2.5 2.5 0 0 0 4.5 14v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14a2.5 2.5 0 0 0 2.5-2.5h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14A2.5 2.5 0 0 0 11.5 2V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5zm1 4.5h3A1.5 1.5 0 0 1 11 6.5v3A1.5 1.5 0 0 1 9.5 11h-3A1.5 1.5 0 0 1 5 9.5v-3A1.5 1.5 0 0 1 6.5 5z"/>
				</symbol>
				<symbol id="geo-fill" viewBox="0 0 16 16">
					<path fill-rule="evenodd" d="M4 4a4 4 0 1 1 4.5 3.969V13.5a.5.5 0 0 1-1 0V7.97A4 4 0 0 1 4 3.999zm2.493 8.574a.5.5 0 0 1-.411.575c-.712.118-1.28.295-1.655.493a1.319 1.319 0 0 0-.37.265.301.301 0 0 0-.057.09V14l.002.008a.147.147 0 0 0 .016.033.617.617 0 0 0 .145.15c.165.13.435.27.813.395.751.25 1.82.414 3.024.414s2.273-.163 3.024-.414c.378-.126.648-.265.813-.395a.619.619 0 0 0 .146-.15.148.148 0 0 0 .015-.033L12 14v-.004a.301.301 0 0 0-.057-.09 1.318 1.318 0 0 0-.37-.264c-.376-.198-.943-.375-1.655-.493a.5.5 0 1 1 .164-.986c.77.127 1.452.328 1.957.594C12.5 13 13 13.4 13 14c0 .426-.26.752-.544.977-.29.228-.68.413-1.116.558-.878.293-2.059.465-3.34.465-1.281 0-2.462-.172-3.34-.465-.436-.145-.826-.33-1.116-.558C3.26 14.752 3 14.426 3 14c0-.599.5-1 .961-1.243.505-.266 1.187-.467 1.957-.594a.5.5 0 0 1 .575.411z"/>
				</symbol>
				<symbol id="home" viewBox="0 0 16 16">
				<path d="M8.354 1.146a.5.5 0 0 0-.708 0l-6 6A.5.5 0 0 0 1.5 7.5v7a.5.5 0 0 0 .5.5h4.5a.5.5 0 0 0 .5-.5v-4h2v4a.5.5 0 0 0 .5.5H14a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.146-.354L13 5.793V2.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.293L8.354 1.146zM2.5 14V7.707l5.5-5.5 5.5 5.5V14H10v-4a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v4H2.5z"/>
				</symbol>
			</svg>
		</head>
		<body>
			<div class="container px-4 py-5" id="icon-grid">
				<h2 class="pb-2 border-bottom">Demo</h2>
				<div class="col d-flex align-items-start">
					<svg class="bi text-muted flex-shrink-0 me-3" width="1.75em" height="1.75em"><use xlink:href="#home"/></svg>
					<div>
					<h3 class="fw-bold mb-0 fs-4">Github</h3>
					<br/>
					<p> * Github URL : <a href="https://github.com/hellices/shcWasmGo">https://github.com/hellices/shcWasmGo</a></p>
					</div>
				</div>
				<br/>
				<div class="col d-flex align-items-start">
					<svg class="bi text-muted flex-shrink-0 me-3" width="1.75em" height="1.75em"><use xlink:href="#tools"/></svg>
					<div>
						<h3 class="fw-bold mb-0 fs-4">Pre-requests</h3>
						<br/>
						<p> * golang install : <a href="javascript:void(0);">https://go.dev/doc/install</a></p>
						<p> * tinygo install : <a href="javascript:void(0);">https://tinygo.org/getting-started/install/macos/</a></p>
						<p> * spin install : <a href="javascript:void(0);">curl -fsSL https://developer.fermyon.com/downloads/install.sh | bash</a></p>
					</div>
				</div>
				<br/>
				<div class="col d-flex align-items-start">
					<svg class="bi text-muted flex-shrink-0 me-3" width="1.75em" height="1.75em"><use xlink:href="#cpu-fill"/></svg>
					<div>
					<h3 class="fw-bold mb-0 fs-4">Build Example</h3>
					<br/>
					<p> * Command : <b>tinygo build -target=wasi -gc=leaking -no-debug -o example.wasm main.go</b></p>
					<p> * Create example.wasm file </p>
					</div>
				</div>
				<br/>
				<div class="col d-flex align-items-start">
				<svg class="bi text-muted flex-shrink-0 me-3" width="1.75em" height="1.75em"><use xlink:href="#geo-fill"/></svg>
				<div>
					<h3 class="fw-bold mb-0 fs-4">Run with Spin FrameWork in Runtime WASI</h3>
					<br/>
					<p> * Command : <b>spin up --listen localhost:3000</b></p>
				</div>
				</div>
			</div>
		</body>
	</html>
	`
)

func GetWasm() string {
	return wasm
}