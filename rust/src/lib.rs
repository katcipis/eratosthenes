struct Primes<I: Iterator> {
    iter: I
}

impl<I: Iterator> Primes<I> {
    fn new(iter: I) -> Primes<I> {
        Primes { iter }
    }
}

impl<I: Iterator> Iterator for Primes<I>
{
    type Item = I::Item;

    fn next(&mut self) -> Option<I::Item> {
        self.iter.next()
    }
}

struct NaturalNumbers {
    cur: u32,
}

impl NaturalNumbers {
    fn new(start: u32) -> NaturalNumbers {
        NaturalNumbers { cur: start }
    }
}

impl Iterator for NaturalNumbers {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        // TODO: handle overflow, won't panic in non-debug mode
        let val = self.cur;
        self.cur += 1;
        Some(val)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn natural_numbers_generation() {
        let mut nums = NaturalNumbers::new(2);

        assert_eq!(nums.next(), Some(2));
        assert_eq!(nums.next(), Some(3));
        assert_eq!(nums.next(), Some(4));
        assert_eq!(nums.next(), Some(5));

        let mut nums = NaturalNumbers::new(100);

        assert_eq!(nums.next(), Some(100));
        assert_eq!(nums.next(), Some(101));
        assert_eq!(nums.next(), Some(102));
    }

    //#[test]
    //fn primes_generation() {
        //let mut primes = Primes::new();

        //assert_eq!(primes.next(), Some(2));
        //assert_eq!(primes.next(), Some(3));
        //assert_eq!(primes.next(), Some(5));
        //assert_eq!(primes.next(), Some(7));
        //assert_eq!(primes.next(), Some(11));
        //assert_eq!(primes.next(), Some(13));
        //assert_eq!(primes.next(), Some(17));
        //assert_eq!(primes.next(), Some(19));
        //assert_eq!(primes.next(), Some(23));
        //assert_eq!(primes.next(), Some(29));
    //}
}
