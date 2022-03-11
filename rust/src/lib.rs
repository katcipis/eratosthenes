pub fn primes(n: u32) -> Vec<u32> {
    let mut vec: Vec<u32> = Vec::new();
    // Did lose a few hours with how to have an iterator that incrementally
    // chains more and more filters. This is the only thing that worked:
    //
    // - https://users.rust-lang.org/t/multiple-staggered-filters-on-an-iterator/60959/4
    //
    // The learning curve for this stuff is steep as fuck x_x.
    // Tried to model primes also as an iterator/stream, but the borrow
    // checker defeated me at my current novice level x_x.
    // Code very similar to this, but using a struct with the boxed type,
    // simply didn't work, there was always some different lifetime/reference
    // issue that I was unable to solve properly x_x.
    // In the end it could have worked if it was easy to declare a trait that is
    // an Iterator + Clone, but it is not easy AFAIK.
    let mut iter: Box<dyn Iterator<Item = u32> + '_> = Box::new(NaturalNumbers::new(2));

    for _ in 0..n {
        let val = iter.next().expect("unexpected EOS");
        vec.push(val);
        iter = Box::new(iter.filter(move |x| x % val != 0));
    }

    vec
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

    #[test]
    fn primes_generation() {
        let p = primes(0);

        assert_eq!(p.len(), 0);

        let p = primes(1);

        assert_eq!(p.len(), 1);
        assert_eq!(p[0], 2);

        let p = primes(10);

        assert_eq!(p.len(), 10);
        assert_eq!(p[0], 2);
        assert_eq!(p[1], 3);
        assert_eq!(p[2], 5);
        assert_eq!(p[3], 7);
        assert_eq!(p[4], 11);
        assert_eq!(p[5], 13);
        assert_eq!(p[6], 17);
        assert_eq!(p[7], 19);
        assert_eq!(p[8], 23);
        assert_eq!(p[9], 29);
    }
}
