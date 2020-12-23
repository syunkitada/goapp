function escape_id(str: string): string {
    str = str.replace("/", "-");
    return str;
}

const index = {
    escape_id
};
export default index;
