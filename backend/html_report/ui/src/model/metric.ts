export interface Metric {
    label: String,
    value: number,
    explanation: String,
    operation: String | null,
    submetric: Submetric[],
    rating: number,
}


export interface Submetric {
    label: String,
    value: number,
    explanation: String,
    operation: String | null,
    order: number,
}


export let dummyMetric: Metric = {
    label: "",
    value: -1,
    explanation: "",
    operation: null,
    submetric: [],
    rating: -1,
}

export let dummySubmetric: Submetric = {
    label: "asd",
    value: 0,
    explanation: "",
    operation: "",
    order: 0,
}

export const evaluateMetric = "evaluate-metric";
export const addParenthesis = "add-parenthesis";
export const createMetric = "create-metric";

export enum MetricSteps {
    EvaluateMetric = evaluateMetric,
    AddParenthesis = addParenthesis,
    CreateMetric = createMetric
}  

export enum AddMetricSteps {
    AddMetric = "add-metric",
    ChooseMethod = "choose-method",
    AddingMetric = "adding-metric"
}

export const fromDocument = "From Document"
export const setManually = "Set Manually"
export const fromOutsideSource = "From Outside Source"

export enum CreateMetricOptions {
    FromDocument = fromDocument,
    SetManually = setManually,
    FromOutsideSource = fromOutsideSource
}

export const label = "setting-label"
export const value = "setting-value"
export const explanation = "setting-explanation"

export enum LabelValueSteps {
    Label = label,
    Value = value,
    Explanation = explanation,
}