import React, { Fragment } from 'react';

const CardList = ({list}) => {
    return (
        <Fragment>
            {
                Object.values(list).map((value, key) => {
                    return (<div>Item</div>)
                })
            }
        </Fragment>
    )
}

export default CardList;