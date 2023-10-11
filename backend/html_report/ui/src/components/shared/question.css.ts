import { css } from "lit";


export default css`

:host {

    border-radius: 50%;
    background-color: var(--gray80);
    color: var(--gray20);
    font-weight: 700;
    padding: var(--spacing);
    width: 16px;
    height: 16px;
    display: flex;
    justify-content: center;
    font-size: 1.6rem;
    align-items: center;
    position: relative;
}

span {
    display: none;
    transform: translateX(100%);
    right: 0;
    position: absolute;
    background-color: var(--gray30);
    padding: var(--spacingQuarter);
    color: var(--gray80);
    font-weight: 400;
    border-radius: 10px;
}

:host(:hover) span {
    display: block;
}


`;