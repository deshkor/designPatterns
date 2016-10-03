public class WeakEnemy implements Enemy {

	WeakEnemy() {
		System.out.println("A weak enemy has been born !!");
	}

	private int hp = 10, damage = 5, armor = 3;

	@Override
	public int attack() {
		return this.damage;
	}

	@Override
	public void defend(int damage) {
		if (this.armor >= damage) {
			return;
		}

		this.hp -= (damage - armor);
	}

	@Override
	public int getHp() {
		return this.hp;
	}
}
