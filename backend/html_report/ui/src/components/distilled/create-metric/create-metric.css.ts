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

button, input {
    font-family: NeueMachina !important;
      
}

figure {
    padding: var(--spacingHalf) 9.8rem;
}

.configure-metric-div {
    position: absolute;
    display: flex;
    justify-content: start;
    align-items: end;
    flex-direction: column;
    gap: 1.2rem;
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
    margin-bottom: var(--spacing)
}

.create-metric,
.add-another-metric,
.remove-another-metric {
    color: var(--gray80);
    background: none;
    box-shadow: 0 0 0 1px var(--gray80);
    font-size: 2rem;
    cursor: pointer;  
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    padding: var(--spacingHalf) var(--spacing);
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
    color: var(--gray92);
    box-shadow: 0 0 0 1px var(--gray92);
}

.create-metric {
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    position: absolute;
    left: 0;
    transform: translate(-60%);
    top: 20rem;
}

.remove-another-metric {
    top: 8rem;
}

label {
    font-size: 2rem;
}

.create-metric-button:hover {
    border: 2px solid var(--gray80);
    color: var(-gray80);
}

input {
    padding: var(--spacingHalf);
    border-radius: 10px;
    font-size: 2rem;
    background-color: transparent;
    border: none;
    box-shadow: 0 0 0 1px var(--gray92);
    color: var(--gray92);
}

label {
    display: flex;
    gap: .6rem;
    align-items: center;
}

input::placeholder {
    color: var(--gray60);
}

.operation-div {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: var(--spacingQuarter) 0;
}

.metric-div {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 2.4rem;
    background-color: var(--gray25);
    padding: var(--spacing);
    border-radius: 10px;
    position: relative;
}

.labelinput {
    flex-direction: column;
    display: flex;
    gap: .6rem;
    color: var(--gray80);
}

.add-metric {
    opacity: 0;
    visibility: 0;
    width: 100%;
    height: 100%;
    background-color: transparent;
    display: flex;
    position: absolute;
    font-size: 3rem;
    color: var(--gray80);
    justify-content: center;
    align-items: center;
    -webkit-backdrop-filter: blur(0.075em);
    backdrop-filter: blur(0.075em);
    transition: all .2s;
    border: 1px solid var(--info);
    border-radius: 10px;
    gap: 1.2rem;
    
}

.metric-div {
    z-index: 1;
}

.metric-div:hover .add-metric {
    opacity: 1;
    visibility: 1;
    z-index: 1;

    
}
.add-metric span {
    transition: all .5s;
    opacity: 0;
    padding-left: var(--spacingQuarter);
}

.metric-div:hover span {
    transform: translateY(-2px);
    opacity: 1;
}

.add-metric button {
    background-color: var(--gray30);
    padding: var(--spacingHalf) var(--spacing);
    border-radius: 10px;
    color: var(--gray80);
    font-size: 2rem;
    border: none;
}

.add-metric button {
cursor: pointer;
transition: all .2s;
// border: 1px solid var(--secondary-3);
box-shadow: 0 0 0 0 var(--infoColorD10);
}

.add-metric button:hover {
    box-shadow: 0 0 0 1px var(--infoColorD10);
}

.finish-editing {
    position: absolute;
    right: 5%;
    background: transparent;
    padding: var(--spacingHalf);
    border: 0;
    box-shadow: 0 0 0 1px var(--info);
    cursor: pointer;
    font-size: 2rem;
    color: var(--info);
    border-radius: 10px;
}
 
`;