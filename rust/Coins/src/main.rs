//https://www.spoj.com/problems/COINS/
pub struct Coin {
    value: i32,
}

impl Coin {
    pub fn new(value: i32) -> Self {
        Coin {
            value,
        }
    }

    fn fractional(&self) -> i32 {
        self.value / 2 + self.value / 3 + self.value / 4
    }

    pub fn max_value(&self) -> i32 {
        let frac = self.fractional();
        if frac > self.value {
            return frac
        }
        self.value
    }
}

fn main() {
    let coin_twelve = Coin::new(12);
    let coin_two = Coin::new(2);
    println!("{}", coin_twelve.max_value());
    println!("{}", coin_two.max_value());
}

#[cfg(test)]
mod test {
    use Coin;

    #[test]
    fn twelve() {
        let expected = 13;
        let actual   = Coin::new(12).max_value();
        assert_eq!(expected, actual);
    }

    #[test]
    fn two() {
        let expected = 2;
        let actual   = Coin::new(2).max_value();
        assert_eq!(expected, actual);
    }
}
