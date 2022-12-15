import axios from "axios";


axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem("identifier");
const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

export default instance;
