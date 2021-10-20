// written after docs 19/10-21
// https://gitlab.pcvdata.dk/super-team-euxtra/neocheckin/docs

import express from "express";
import cors from "cors";
import { readFile } from "fs/promises"
import { join } from "path";

const exists = (...args) => {
    for (let i in args) {
        if (args[i] === undefined || args[i] === null) {
            return false
        }
    }
    return true
}


interface TaskType {
    id: number,
    name: string,
    description: string,
    active: boolean,
    schedule: {
        from: {
            hour: number,
            minute: number,
            second: number,
        },
        to: {
            hour: number,
            minute: number,
            second: number,
        },
        days: {
            monday: boolean,
            tuesday: boolean,
            wednesday: boolean,
            thursday: boolean,
            friday: boolean,
            saturday: boolean,
            sunday: boolean,
        }
    },
}

const tasks: TaskType[] = [
    {
        id: 0,
        name: "name0",
        description: "desc0",
        active: true,
        schedule: {
            from: {
                hour: 0,
                minute: 0,
                second: 0,
            },
            to: {
                hour: 24,
                minute: 0,
                second: 0,
            },
            days: {
                monday: true,
                tuesday: true,
                wednesday: true,
                thursday: true,
                friday: true,
                saturday: false,
                sunday: false,
            }
        },
    },
    {
        id: 1,
        name: "name1",
        description: "desc1",
        active: true,
        schedule: {
            from: {
                hour: 0,
                minute: 0,
                second: 0,
            },
            to: {
                hour: 24,
                minute: 0,
                second: 0,
            },
            days: {
                monday: true,
                tuesday: true,
                wednesday: true,
                thursday: true,
                friday: true,
                saturday: false,
                sunday: false,
            }
        },
    },
];

interface addTaskRequest {
    taskId: number,
    name: string,
    employeeId: string,
    highLevelApiKey: string,
    systemIdentifier: string,
}


const server = () => {
    const app = express();
    app.use(express.json());
    app.use(cors());

    app.get('/api/tasks/types', (req, res) => {
        return res.status(200).json({
            data: tasks, 
        });
    });
    
    app.post('/api/tasks/add', (req, res) => {
        const {taskId, name, employeeId, highLevelApiKey, systemIdentifier}: addTaskRequest = req.body;
        if (exists(taskId, name, employeeId, highLevelApiKey, systemIdentifier)) {

        } else {
            return res.status(400).json({error: "missing values"})
        }

        return res.status(200).json();
    });

    app.listen(7000, () => {
        console.log('server started')
    })
}

server();