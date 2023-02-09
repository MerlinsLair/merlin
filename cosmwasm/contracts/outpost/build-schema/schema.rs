use cosmwasm_schema::write_api;

use outpost::msg::{ExecuteMsg, InstantiateMsg, QueryMsg};

fn main() {
    write_api! {
        name: "merlin-outpost",
        instantiate: InstantiateMsg,
        query: QueryMsg,
        execute: ExecuteMsg,
    };
}
