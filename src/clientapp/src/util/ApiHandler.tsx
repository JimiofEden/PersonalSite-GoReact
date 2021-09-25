import axios from 'axios';

// TODO - URLs should be defined somewhere in an app setting
export const fetchSkills = () => {
    return axios.post(
        `http://localhost:9876/api/data`, 
        {
            query: `
                query {
                    skills {
                        id
                        skillName
                        skillType {
                            id
                            skillTypeId
                            skillTypeName
                            sequence
                        }
                        url
                        comment
                        sequence
                    }
                }
            `
        }
        , {
        headers: {
          'Content-Type': 'application/json'
        }
    });
}

export const fetchSkillTypes = () => {
    return axios.post(
        `http://localhost:9876/api/data`, 
        {
            query: `
                query {
                    skillTypes {
                        id
                        skillTypeId
                        skillTypeName
                        sequence
                    }
                }
            `
        }
        , {
        headers: {
          'Content-Type': 'application/json'
        }
    });
}

export const fetchLinks = () => {
    return axios.get(`http://localhost:9876/api/links`);
}