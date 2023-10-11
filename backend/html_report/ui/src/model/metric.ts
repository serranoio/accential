export interface Metric {
    id: number,
    documentId: number,
    label: String,
    value: number,
    explanation: String,
    operation: String | null,
    submetric: Submetric[],
    rating: string,
}


export interface Submetric {
    id: number
    label: String,
    value: number,
    explanation: String,
    operation: String | null,
    order: number,
}


export let dummyMetric: Metric = {
    id: 0,
    documentId: 0,
    label: "",
    value: -1,
    explanation: "",
    operation: null,
    submetric: [],
    rating: "",
}

export let dummySubmetric: Submetric = {
    id: 0,
    label: "",
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
export const fromOthers = "From Others"


export enum CreateMetricOptions {
    FromDocument = fromDocument,
    SetManually = setManually,
    FromOutsideSource = fromOutsideSource,
    FromOthers = fromOthers
}

export const label = "setting-label"
export const value = "setting-value"
export const explanation = "setting-explanation"

export enum LabelValueSteps {
    Label = label,
    Value = value,
    Explanation = explanation,
}