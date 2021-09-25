import axios from 'axios';
import { loader } from 'graphql.macro';
import { print } from 'graphql/language/printer';


// TODO - URLs should be defined somewhere in an app setting
export const fetchSkills = () => {
    const queryGql = loader('./queries/skills.gql');
    const query = print(queryGql)
    return axios.post(
        `http://localhost:9876/api/data`, 
        { query: query },
        { headers: { 'Content-Type': 'application/json' } }
    );
}

export const fetchSkillTypes = () => {
    const queryGql = loader('./queries/skillTypes.gql');
    const query = print(queryGql)
    return axios.post(
        `http://localhost:9876/api/data`, 
        { query: query },
        { headers: { 'Content-Type': 'application/json' } }
    );
}

export const fetchLinks = () => {
    const queryGql = loader('./queries/storedLinks.gql');
    const query = print(queryGql)
    return axios.post(
        `http://localhost:9876/api/data`, 
        { query: query },
        { headers: { 'Content-Type': 'application/json' } }
    );
}