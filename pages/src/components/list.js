import React, { useState, useEffect } from "react";
import axios from "axios";

export default function List() {
	const [files, setFiles] = useState([]);
	const [checked, setChecked] = useState([]);

	const getColistFiles = () => {
		axios.get("/list").then((resp) => {
		setFiles(resp.data);
		});
	};

	useEffect(() => {
		getColistFiles();
	}, [])

	const handleCheck = (event) => {
		var updatedList = [...checked];
		if (event.target.checked) {
			updatedList = [...checked, event.target.value];
		} else {
			updatedList.splice(checked.indexOf(event.target.value), 1);
		}
		setChecked(updatedList);
		console.log(updatedList)
	};

	const onDelete= () => {
		axios.post("/remove", checked);
	}

	if (files.length > 0) {
  		return (
			<div>
			{ 
				files.map((file, index) => (
				<div key={index}>
					<input value={file} type="checkbox" onChange={handleCheck}/>
					<span>{file}</span>
				</div>
				))
			}
				<button onClick={onDelete}>
					Delete files
				</button>
			</div>
		);
		} else {
		return (
			<h4> No files in file store </h4>
		);
		}
}

