import axios from "axios";


// axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.identifier;
const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});
instance.interceptors.request.use(
	function(config) {
	  const token = localStorage.getItem("identifier"); 
	  if (token) {
		config.headers["Authorization"] = 'Bearer ' + token;
	  }
	  return config;
	},
	function(error) {
	  return Promise.reject(error);
	}
  );

export default instance;
