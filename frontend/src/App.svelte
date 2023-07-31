<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {
		GetPortsList,
		OpenPort,
		ClosePort,
		Send } from '../wailsjs/go/main/App.js'

	// 阿付看这里, 这里是事件监听
	import { EventsOn } from '../wailsjs/runtime'

  let resultText: string = "Please enter your name below 👇"
  let name: string
	let timeText: string = "The time is: "
	let openStatus: string = "None"
	let readData: string = "None"

  function getPorts(): void {
    GetPortsList().then(
			result => resultText = result
		)
  }

	function openPort_1(): void {
    OpenPort("/dev/ttyUSB0", 9600).then(result => openStatus = result)
	}

	function closePort_1(): void {
		ClosePort().then(result => openStatus = result)
	}

	function send(): void {
		let data = "Hello World"
		// data to byte array
		// let dataBytes = new TextEncoder().encode(data)
		// Send(dataBytes).then(result => resultText = result)
		
		// data to uint8array
		// let dataBytes = new TextEncoder().encode(data)
		// let dataUint8Array = new Uint8Array(dataBytes)
		// console.log(dataUint8Array)
		// Send(dataUint8Array).then(result => resultText = result)

		// Send(dataUint8Array).then(result => resultText = result)

		// data to base64
		let dataBase64 = btoa(data)
		Send(dataBase64).then(result => resultText = result)



	}

	// read data from serial port
	EventsOn('read', (base64data)=>{
		// console.log(res)
		base64data = base64data.replace(/[\r\n]/g,"")
		let data = atob(base64data)
		readData = data

	}, -1)

	// 阿付看这里, 这里是事件监听
	EventsOn('test', (res)=>{
		// console.log(res)
		timeText = res
	}, -1)

	// 串口异常已被关闭
	EventsOn('error', (res)=>{
		// console.log(res)
		openStatus = res
	}, -1)
</script>

<main>
  <div class="result" id="readdata">{readData}</div>
  <div class="result" id="time">{timeText}</div>
  <div class="result" id="openstatus">{openStatus}</div>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <button on:click={getPorts}>获取串口列表</button>
    <button on:click={openPort_1}>开关串口一</button>
    <button on:click={closePort_1}>关闭串口一</button>
		<button on:click={send}>发送</button>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
