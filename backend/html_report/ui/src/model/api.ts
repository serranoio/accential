import { Metric } from "./metric"

const Metrics = "metric"
const Get = "get"
// const Update = "update"
const Add = "add"
const Delete = "delete"
// const All = "all"
const Api = "api"
const Document = "document"
const ChangeName = "change-name"
const GetName = "get-name"
const Post = "post"

const url = 'http://127.0.0.1:8080'
// const url = 'https://accential.fly.dev'

export const GetAllMetrics = async (id: string) => {
    try {
        const metrics = await fetch(`${url}/${Api}/${Metrics}/${Get}/${id}`)

        const json = await metrics.json()

        return json.message
    } catch (err) {
        console.log(err)
    }
}

export const DeleteMetricDb = async (metric: Metric, docId: string) => {

    try {
        const metrics = await fetch(`${url}/${Api}/${Metrics}/${Delete}/${docId}`, {
            method: "POST",
            body: JSON.stringify(
                metric
                ),
        }
        )

        const json = await metrics.json()

        console.log(json)

        return json.message
    } catch (err) {
        console.log(err)
    }

}
export const AddNewMetric = async (newMetric: Metric, docId: string) => {

    try {
        const metrics = await fetch(`${url}/${Api}/${Metrics}/${Add}/${docId}`, {
            method: "POST",
            body: JSON.stringify(
                newMetric
                ),
        }
        )

        const json = await metrics.json()

        console.log(json)

        return json.message
    } catch (err) {
        console.log(err)
    }

}

export const ChangeDocumentName = async (id: string, newName: string) => {
    try {
        const metrics = await fetch(`${url}/${Api}/${Post}/${Document}/${ChangeName}`, {
            method: "POST",
            body: JSON.stringify(
                {
                    id: id, 
                    Name: newName,
                }
            ),
        }
        )

        const json = await metrics.json()

        console.log(json)

        return json.message
    } catch(err) {

    }
}

export const FetchName = async (id: string) => {

    try {
        const metrics = await fetch(`${url}/${Api}/${Post}/${Document}/${GetName}`, 
            {
                method: "POST",
                body: JSON.stringify(
                     {
                        id: id, 
                    }
                ),
            }
        )

        const json = await metrics.json()

    
        return json.message
    } catch(err) {

    }
    
}