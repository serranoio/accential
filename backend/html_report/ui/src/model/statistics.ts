interface MeasureOfWorkingCapital {
    Years: string[],
    TotalAssets: number[],
    TotalLiabilities: number[],
    WorkingCapital:   number[],
    Info:             string,
}

const dummyMowc: MeasureOfWorkingCapital = {
    Years: ["2021", "2020"],
    TotalAssets: [1404235000,
        1372987000
    ],
    TotalLiabilities: [
        1353729000,
        1335001000
       ],
    WorkingCapital: [
        1.0373087966646204,
        1.0284539112704785
    ],
    Info: "The Measure of Working Capital is the Ratio of Total Assets / Total Liablities",
}

export interface Statistics {
    Mowc: MeasureOfWorkingCapital;
}


export const dummyStatistics: Statistics = {
    Mowc: dummyMowc,
}

interface Doc {
    name: string,
}

export const dummyDoc: Doc = {
    name: "Siebert Corp"
}

