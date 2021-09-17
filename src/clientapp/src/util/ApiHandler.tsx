import axios from 'axios';

// TODO - URLs should be defined somewhere in an app setting
export const fetchSkills = () => {
    return axios.get(`http://localhost:9876/api/skills`);
}

export const fetchLinks = () => {
    return axios.get(`http://localhost:9876/api/links`);
}