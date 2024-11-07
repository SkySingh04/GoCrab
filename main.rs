fn main() {
    let x = 10;
    let y = 20.5;
    let z = x + y - 3 * (5 / 2) % 2;
    let a = !x && y || z;
    let c = x > y && y <= z || x == 10;
    let result = if x != y { "unequal" } else { "equal" };
    let complex = (x << 2) + (y >> 1) - (a ^ c);
}
