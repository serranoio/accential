import { css } from "lit";


export default css`

:host {
    position: relative;
    min-height: calc(100vh - 80px);
    display: block;
    background-color: var(--gray22);
    border-top-right-radius: 10px;
    border-top-left-radius: 10px;
}
* {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
}

figure {
    padding: 2.4rem 9.8rem;
}

.configure-metric-div {
    position: absolute;
    display: flex;
    justify-content: start;
    align-items: end;
    flex-direction: column;
    gap: 2.4rem;
    right: 0;
    transform: translateX(50%);
}

h2 {
    color: var(--gray80);
    font-size: 6.4rem;
}
p {
    color: var(--gray60);
    font-size: 2rem;
    margin-bottom: 2.4rem;
}

.create-metric,
.add-another-metric,
.remove-another-metric {
    color: var(--gray60);
    background: none;
    box-shadow: 0 0 0 1px var(--gray60);
    font-size: 2rem;
    cursor: pointer;  
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    padding: 1.2rem 3.6rem;
    cursor: pointer;
    transition: all .2s;
}

.add-another-metric,
.remove-another-metric {
    border-top-right-radius: 10px;
    border-bottom-right-radius: 10px;
}

.create-metric:hover,
.add-another-metric:hover,
.remove-another-metric:hover {
    color: var(--gray80);
    box-shadow: 0 0 0 1px var(--gray80);
}

.create-metric {
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    position: absolute;
    left: 0;
    transform: translate(-50%);
    top: 20rem;
}

.remove-another-metric {
    top: 8rem;
}



.create-metric-button:hover {
    border: 2px solid var(--gray80);
    color: var(-gray80);
}



input {
    padding: 1.2rem;
    border-radius: 10px;
    font-size: 2rem;
    background-color: transparent;
    border: none;
    box-shadow: 0 0 0 1px var(--gray80);
    color: var(--gray80);
}
.operation-div {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2.4rem 0;
}

.metric-div {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: .6rem;
}

 
`;