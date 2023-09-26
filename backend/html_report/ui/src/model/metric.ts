export interface Metric {
    Label: String,
    Value: number,
    Explanation: String,
    Operation: String | null,
    NextMetric: Metric | null,
}

export const dummyMetric: Metric = {
    Label: "",
    Value: -1,
    Explanation: "",
    Operation: null,
    NextMetric: null,
}

export const evaluateMetric = "evaluate-metric";
export const addParenthesis = "add-parenthesis";
export const createMetric = "create-metric";


export enum MetricSteps {
    EvaluateMetric = evaluateMetric,
    AddParenthesis = addParenthesis,
    CreateMetric = createMetric
}  

