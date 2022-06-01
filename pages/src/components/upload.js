import React, { useState, useEffect } from "react";
import axios from 'axios';

export default function Upload() {
	const [file, setFile] = useState([]);

	const onFileChange = event => {
		console.log(event)
		console.log(event.target.files[0])
		setFile(event.target.files[0])
		console.log(file)
	};

	const onFileUpload = () => {
		const formData = new FormData();
		console.log(file);
		formData.append(
			"File",
			file,
			file.name
		);
		axios.post("/add", formData);
	};

	const fileData = () => {
		return (
			<div>
			<h2>File Details:</h2>
			<p> File Name: {file.name} </p>
			</div>
		);
	};

	return (
		<div>
		<div>
		<input type="file" onChange={onFileChange} />
		<button onClick={onFileUpload}>
		Upload!
		</button>
		</div>
		{fileData()}
		</div>
	);
}
