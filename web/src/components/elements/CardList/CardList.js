import React, { Fragment } from "react";

const CardList = ({ list }) => {
  return (
    <Fragment>
      {Object.values(list).map(() => {
        return <div>Item</div>;
      })}
    </Fragment>
  );
};

export default CardList;
