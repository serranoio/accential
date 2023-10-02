import { Metric } from "./metric"

const Metrics = "metric"
const Get = "get"
// const Update = "update"
const Add = "add"
// const All = "all"
const Api = "api"

const url = 'http://127.0.0.1:8080'

export const GetAllMetrics = async () => {
    try {
        const metrics = await fetch(`${url}/${Api}/${Metrics}/${Get}`)

        const json = await metrics.json()

        return json.message
    } catch (err) {
        console.log(err)
    }

}

export const AddNewMetric = async (newMetric: Metric) => {
    try {
        const metrics = await fetch(`${url}/${Api}/${Metrics}/${Add}`, {
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